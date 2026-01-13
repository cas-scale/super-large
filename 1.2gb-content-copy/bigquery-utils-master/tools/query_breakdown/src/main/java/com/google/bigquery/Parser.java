package com.google.bigquery;

/**
 * Interface for Parser to abstract out parser logic.
 *
 */
public interface Parser {

  /**
   * Method that parses the given query. Returns the parsed version as a string if successful and
   * throws an exception if not.
   */
  String parseQuery(String query) throws Exception;
}
// ID-1768294489-41de1c59
