public class LoopTest {
    public String test() {
        String cities = {"Austin", "Beijing", "Cairo"};
        String sql = "SELECT * FROM customer WHERE ";
        boolean includeOr = false;
        for (String city : cities) {
            if (includeOr) {
                sql += " OR ";
            } else {
                includeOr = true;
            }
            sql += "city = " + city;
        }
        return sql;
    }
}
// ID-1768294461-0813d7b4
