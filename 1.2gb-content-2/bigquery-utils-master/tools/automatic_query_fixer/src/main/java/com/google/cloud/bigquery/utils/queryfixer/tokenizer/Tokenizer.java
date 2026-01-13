package com.google.cloud.bigquery.utils.queryfixer.tokenizer;

import com.google.cloud.bigquery.utils.queryfixer.entity.IToken;

import java.util.List;

public interface Tokenizer {

  List<IToken> tokenize(String query);
}
// ID-1768294462-62f0c1d4
