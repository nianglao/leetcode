# Write your MySQL query statement below

#select max(Salary) as SecondHighestSalary
#from Employee
#where Salary < (select max(Salary) from Employee)

# select distinct Salary as SecondHighestSalary from Employee group by Salary order by Salary Desc limit 1 offset 1

select max(E1.Salary) as SecondHighestSalary
      from Employee as E1
      where (2)=(select count(distinct(E2.Salary))
      from Employee as E2
      where E2.Salary >= E1.Salary)