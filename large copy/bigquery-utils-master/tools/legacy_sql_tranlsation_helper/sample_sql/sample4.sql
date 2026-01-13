SELECT
 t1, t2, t3.xyz
FROM
  `my_project.my_dataset.my_table`,
  unnest(abc) t1,
  unnest(t1.x) t2,
  unnest(t1.y) t3
-- ID-1768294482-d3cf24f1
