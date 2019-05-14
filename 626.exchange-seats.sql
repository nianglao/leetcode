# Write your MySQL query statement below

select s1.id,coalesce(s2.student,s1.student) as student
from seat s1
left join seat s2
on ((s1.id)=s2.id-1 and s1.id %2!=0 or (s1.id=s2.id+1 and s1.id%2=0))
order by 1