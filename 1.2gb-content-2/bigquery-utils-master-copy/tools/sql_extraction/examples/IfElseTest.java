public class IfElseTest {
    public String test(int i) {
        String sql = "SELECT * FROM customer";
        if (i > 0) {
            sql += " WHERE id = " + i;
        }
        return sql;
    }
}
// ID-1768294461-2d51b4d0
