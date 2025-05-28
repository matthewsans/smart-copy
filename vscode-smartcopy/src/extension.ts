import * as vscode from 'vscode';
import * as cp from 'child_process';
import * as path from 'path';
import { spawn } from 'child_process';

async function copyOpenEditors() {
    const editors = vscode.workspace.textDocuments;
    const wsFolders = vscode.workspace.workspaceFolders ?? [];
    var count = 0;
    let payload = '';

    if (editors.length === 0) {
      vscode.window.showInformationMessage('No open editors to copy.'); 
      return;
    }

    for (const ed of editors) {
      if  (ed.uri.scheme !== 'file' ||                            // skip virtual docs (git, gitlens, debug, etc.)
          (ed.fileName.includes(`${path.sep}.git${path.sep}`)) ||  // skip .git internals
          (ed.fileName.endsWith('.git'))) continue;

      const headerPath = vscode.workspace.asRelativePath(ed.uri, false);
      payload += `\n// ---- ${headerPath} ----\n\n`;
      payload += ed.getText() + '\n';
      count = count + 1;
    }


  
    await vscode.env.clipboard.writeText(payload);
    vscode.window.showInformationMessage(
      `✔ Copied ${count} open editor(s) from ${wsFolders.length || 1} workspace folder(s).`
    );
          

  }

export function activate(context: vscode.ExtensionContext) {
  console.log('✅ vscode-smartcopy activated!');

  // Helper to invoke the Go binary
  function runSmartCopy(uri: vscode.Uri, mode: string) {
    const goBinaryName = process.platform === 'win32'
      ? 'smartcopy.exe'
      : 'smartcopy';
    const platformDir = process.platform === 'darwin'
      ? 'mac'
      : process.platform;
    const binaryPath = path.join(
      context.extensionPath, '..', 'bin', platformDir, goBinaryName
    );
    const args = [ uri.fsPath, '--mode', mode ];

    const child = spawn(binaryPath, args, {
      stdio: ['ignore', 'ignore', 'pipe'],   
    });

    let stderr = '';
    child.stderr.on('data', chunk => stderr += chunk);

    child.on('close', code => {
      if (code !== 0) {
        vscode.window.showErrorMessage(`Smart Copy failed: ${stderr.trim()}`);
      } else {
        vscode.window.showInformationMessage(`✔ Smart Copy (${mode}) completed.`);
      }
    });
  }

  //register Open Windows copy
  const dispWin = vscode.commands.registerCommand('vscode-smartcopy.copyOpenEditors', 
    copyOpenEditors
  );
  // Register one command per mode :: Uses GO Scripts
  const modes = [
    { commandId: 'vscode-smartcopy.runMinimal',       mode: 'minimal'   },
    { commandId: 'vscode-smartcopy.runGitignoreList', mode: 'gitignore-list'},
    { commandId: 'vscode-smartcopy.runGitignore',     mode: 'gitignore' },
    { commandId: 'vscode-smartcopy.runAll',           mode: 'all'       }


  ];
  for (const { commandId, mode } of modes) {
    const disp = vscode.commands.registerCommand(commandId, (uri?: vscode.Uri) => {
      if (!uri) {
        const folders = vscode.workspace.workspaceFolders;
        if (folders && folders.length > 0) {
          uri = folders[0].uri; // fallback to workspace root
        } else {
          vscode.window.showErrorMessage('No folder selected and no workspace open.');
          return;
        }
      }
      runSmartCopy(uri, mode);
    });

    context.subscriptions.push(disp);
  }
}

export function deactivate() {}
