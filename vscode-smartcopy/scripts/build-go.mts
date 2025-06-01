import { mkdirSync, existsSync } from 'node:fs';
import { join } from 'node:path';
import { spawnSync } from 'node:child_process';

const targets = [
  ['darwin','arm64'], ['darwin','amd64'],
  ['linux','amd64'],  ['linux','arm64'],
  ['windows','amd64'],
];

const repo   = new URL('../..', import.meta.url).pathname;
const cliDir = repo;                 // or join(repo,'vscode-smartcopy','cli')
const binDir = join(repo,'vscode-smartcopy','bin');

for (const [os,arch] of targets) {
  const outDir = join(binDir, `${os}-${arch}`);
  mkdirSync(outDir, { recursive: true });

  const exeName = os === 'windows' ? 'smartcopy.exe' : 'smartcopy';
  const exePath = join(outDir, exeName);

  console.log(`• building ${os}/${arch}`);
  const { status, error } = spawnSync(
    'go', ['build', '-o', exePath],
    { cwd: cliDir,
      env: { ...process.env, GOOS: os, GOARCH: arch, CGO_ENABLED: '0' },
      stdio: 'inherit',
      shell: false }
  );
  if (status !== 0 || error) throw error ?? new Error('go build failed');
  if (!existsSync(exePath)) throw new Error('binary not created');
}
