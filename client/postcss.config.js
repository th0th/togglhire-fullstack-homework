const path = require('path');

module.exports = {
  plugins: {
    'postcss-nested': {},
    'postcss-flexbugs-fixes': {},
    'postcss-custom-media': {
      importFrom: path.join(__dirname, 'styles', 'breakpoints.css'),
    },
    'postcss-preset-env': {
      autoprefixer: {
        flexbox: 'no-2009',
      },
      stage: 3,
      features: {
        'custom-properties': false,
      },
    },
  },
};
