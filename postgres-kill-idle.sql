select * from pg_catalog.pg_stat_activity 

SELECT pg_terminate_backend(pid)
FROM pg_stat_activity
WHERE datname = 'common'
	AND pid <> pg_backend_pid()
	AND state in ('idle', 'idle in transaction', 'idle in transaction (aborted)', 'disabled') 
	AND state_change < current_timestamp - INTERVAL '30' MINUTE;
