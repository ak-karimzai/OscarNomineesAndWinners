CREATE OR REPLACE FUNCTION get_performance_by_nomination_id(n_id INT)
RETURNS performances AS
$$
DECLARE
  performance performances;
BEGIN
  SELECT p.*
  INTO performance
  FROM performances p
  JOIN nominated_performances np ON np.performance_id = p.id
  WHERE np.nomination_id = n_id;

  RETURN performance;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION get_nominations_by_performance_id(p_id INT)
RETURNS SETOF nominations AS
$$
BEGIN
  RETURN QUERY
  SELECT n.*
  FROM nominations n
  JOIN nominated_performances np ON np.nomination_id = n.id
  WHERE np.performance_id = p_id;
END;
$$ LANGUAGE plpgsql;

-- for testing
-- SELECT * FROM get_nominations_by_performance_id(456);
-- SELECT * FROM get_performance_by_nomination_id(123);