-- name: GetLong :one
select long
from links
where short = $1
limit 1;

-- name: Revoke :exec
delete
from links
where short = $1;

-- name: Create :one
insert into links(short, long)
values ($1, $2)
returning short;
