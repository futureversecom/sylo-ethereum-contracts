module.exports = {
  mocha: {
    grep: "@skip-on-coverage", // Find everything with this tag
    invert: true               // Run the grep's inverse set.
  },
  skipFiles: ['ECDSA.sol', 'Payments/Pricing/Manager.sol', 'Payments/Pricing/Voting.sol']
}