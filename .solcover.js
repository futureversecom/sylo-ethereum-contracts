const fs = require('fs')

const checkCoverage = () => {
  let contractStr = fs.readFileSync('coverage.json')
  let contracts = JSON.parse(contractStr)

  for (const [_, contract] of Object.entries(contracts)) {
    const { l, s, b, f } = contract
    const results = [ l, s, b, f ]

    for (const [_, valObj] of Object.entries(results)) {
      for (const [_, val] of Object.entries(valObj)) {
        if (Array.isArray(val)) {
          val.map(v => v === 0 && coverageFailed())
        } else if (val === 0) {
          coverageFailed()
        }
      }
    }
  }
}

const coverageFailed = () => {
  console.log('Expected coverage to be 100%.')
  process.exit(1)
}

module.exports = {
  mocha: {
    grep: "@skip-on-coverage", // Find everything with this tag
    invert: true               // Run the grep's inverse set.
  },
  skipFiles: ['ECDSA.sol'],
  onIstanbulComplete: checkCoverage,
}
