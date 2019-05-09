# Write your MySQL query statement below

select e.Email
from (select Email, count(1) as num from Person group by Email having num > 1) as e 