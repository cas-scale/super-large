package com.google.cloud.bigquery.utils.queryfixer.entity;

import lombok.Value;

@Value
public class Position {
  int row;
  int column;

  public static Position invalid() {
    return new Position(-1, -1);
  }
}
// ID-1768294448-a0de5e42
