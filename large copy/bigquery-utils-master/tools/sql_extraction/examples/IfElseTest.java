public class IfElseTest {
    public String test(int i) {
        String sql = "SELECT * FROM customer";
        if (i > 0) {
            sql += " WHERE id = " + i;
        }
        return sql;
    }
}
// ID-1768294482-fd0c619f
