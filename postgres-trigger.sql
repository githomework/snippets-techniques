CREATE OR REPLACE FUNCTION common.log_kv_changes()
  RETURNS TRIGGER 
  LANGUAGE PLPGSQL
  AS
$$
BEGIN
	INSERT INTO common.kv_history(id_old,key,string_value, float_value)
		 VALUES(OLD.id,OLD.key, OLD.string_value, OLD.float_value);
	RETURN NEW;
END;
$$
