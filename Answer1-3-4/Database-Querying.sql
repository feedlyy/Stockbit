CREATE TABLE stockbit.Users (
                                Id INT auto_increment NOT NULL primary key,
                                UserName varchar(100) NULL DEFAULT NULL,
                                Parent INT NULL
)
ENGINE=InnoDB
DEFAULT CHARSET=latin1
COLLATE=latin1_swedish_ci;

select x.Id, x.UserName, case when x.Parent = y.Id then y.UserName else NULL end as 'ParentUserName'
from (select * from stockbit.Users u) as x, (select * from stockbit.Users u2) as y
where(x.Parent = y.Id or y.Parent = 0) order by x.UserName limit 3