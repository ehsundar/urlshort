-- name: GetLong :one
select long
from links
where short = $1
limit 1;

-- name: Delete :exec
delete
from links
where short = $1;

-- name: Create :one
insert into links(short, long)
values ($1, $2)
returning short;

-- name: List :many
select short, long, created_at
from links
order by created_at desc;
