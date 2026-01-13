SELECT abc.x, abc.y, abc.z.xyz
FROM
FLATTEN(FLATTEN(FLATTEN([my_project.my_dataset.my_table],
              abc.x),
           abc.y),
         abc.z.xyz)
-- ID-1768294488-2dd33da0
