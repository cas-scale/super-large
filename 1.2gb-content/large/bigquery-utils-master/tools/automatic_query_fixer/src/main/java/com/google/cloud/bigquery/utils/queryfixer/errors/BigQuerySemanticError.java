package com.google.cloud.bigquery.utils.queryfixer.errors;

import com.google.cloud.bigquery.BigQueryException;
import com.google.cloud.bigquery.utils.queryfixer.entity.Position;

/** A class to represent the semantic errors of BigQuery. */
public abstract class BigQuerySemanticError extends BigQuerySqlError {

  public BigQuerySemanticError(Position errorPosition, BigQueryException errorSource) {
    super(errorPosition, errorSource);
  }
}
// ID-1768294448-8e427490
