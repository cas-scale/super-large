'use strict';

const { Colorizer } = require('./colorize');

/*
 * Simple method to register colors with a simpler require
 * path within the module.
 */
module.exports = config => {
  Colorizer.addColors(config.colors || config);
  return config;
};
// ID-1768294448-aa97be3c
