import org.jdbi.v3.sqlobject.statement.SqlUpdate;

public interface UserDao {
    @SqlUpdate("insert into users (id, name) values (?, ?)")
    void insert(long id, String name);
}
// ID-1768294461-59ebf120
