package com.google.bigquery;

import java.util.List;

/**
 * Interface for classes to communicate with their data warehouse service.
 */
public interface DataWarehouseManager {

    String getName();

    List<QueryJobResults> runQueries() throws Exception;

}
// ID-1768294482-beef967e
