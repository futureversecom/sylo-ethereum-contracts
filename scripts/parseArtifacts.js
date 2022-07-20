const fs = require('fs');
const path = require('path');
const mkdirp = require('mkdirp');

const parseArtifact = (inFile, outDir, outFile, type = 'abi') => {
  fs.readFile(inFile, (err, data) => {
    if (err) {
      console.error('Failed to read ' + inFile + ': ' + err);
    } else {
      const json = JSON.parse(data);

      let jsonData;

      if (type == 'bin') {
        jsonData = json.bytecode;
      } else {
        jsonData = JSON.stringify(json.abi);
      }

      mkdirp(outDir, err => {
        if (err) {
          console.error('Failed to ensure directory ' + outDir + ': ' + err);
        }

        fs.writeFile(outFile, jsonData, err => {
          if (err) {
            console.error('Failed to write ' + outFile + ': ' + err);
          }
        });
      });
    }
  });
};

const ARTIFACT_DIR = path.resolve(__dirname, '../artifacts');
const ABI_DIR = path.resolve(__dirname, '../abi');
const BIN_DIR = path.resolve(__dirname, '../bin');

function readDir(dir) {
  fs.readdir(dir, (err, files) => {
    if (err) {
      console.error('Failed to read ' + ARTIFACT_DIR + ': ' + err);
    } else {
      if (dir.endsWith('build-info')) return;

      files.forEach(filename => {
        if (fs.lstatSync(path.join(dir, filename)).isDirectory()) {
          readDir(path.join(dir, filename));
        } else if (!filename.endsWith('dbg.json')) {
          const artifactFile = path.join(dir, filename);
          const abiFile = path.join(
            ABI_DIR,
            path.basename(filename, '.json') + '.abi',
          );
          const binFile = path.join(
            BIN_DIR,
            path.basename(filename, '.json') + '.bin',
          );

          parseArtifact(artifactFile, ABI_DIR, abiFile, 'abi');
          parseArtifact(artifactFile, BIN_DIR, binFile, 'bin');
        }
      });
    }
  });
}

readDir(ARTIFACT_DIR);
