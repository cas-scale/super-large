SELECT max(value$)
FROM `bigquery-public-data.crypto_ethereum.transactions`
group bye hash
LIMIT 10
-- ID-1768294476-af5cedfe
