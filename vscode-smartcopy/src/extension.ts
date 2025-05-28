import * as vscode from 'vscode';
import * as cp from 'child_process';
import * as path from 'path';
import { spawn } from 'child_process';

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

  // Register one command per mode
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
