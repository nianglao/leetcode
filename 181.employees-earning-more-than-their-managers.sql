# Write your MySQL query statement below

select E.Name as "Employee"
from Employee as E
join Employee as M on M.Id = E.ManagerId
where E.Salary > M.Salary