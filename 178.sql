Create table If Not Exists Scores (id int, score DECIMAL(3,2))
Truncate table Scores
insert into Scores (id, score) values ('1', '3.5')
insert into Scores (id, score) values ('2', '3.65')
insert into Scores (id, score) values ('3', '4.0')
insert into Scores (id, score) values ('4', '3.85')
insert into Scores (id, score) values ('5', '4.0')
insert into Scores (id, score) values ('6', '3.65')

/*
SELECT t1.score, @r:=@r+1
FROM ( 
SELECT score 
FROM Scores 
GROUP BY score
ORDER BY score DESC
) t1, (SELECT @r:=0) t2; 
*/



SELECT t3.* FROM Scores AS t4 
LEFT OUTER JOIN (
SELECT t1.score, @r:=@r+1 AS 'rank'
FROM ( 
SELECT score 
FROM Scores 
GROUP BY score
ORDER BY score DESC
) t1, (SELECT @r:=0) t2
) AS t3
ON t4.score = t3.score
ORDER BY t4.score DESC
