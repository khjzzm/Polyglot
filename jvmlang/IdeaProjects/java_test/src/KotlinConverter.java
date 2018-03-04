import java.sql.*;

public class KotlinConverter {
    private static final String DRIVER = "com.mysql.jdbc.Driver";
    private static final String URL = "jdbc:mysql://webdbinstance.cc1j8x3t3gto.ap-northeast-2.rds.amazonaws.com/webdb";
    private static final String USER = "dev";
    private static final String PW = "qrwgDGCGnMPGHtKqQhHnZYmhXfFxHyaJsQuAxnGEv";

    public static void main(String[] args) {
        Connection connection = null;
        Statement st = null;

        try {
            Class.forName(DRIVER);
            connection = DriverManager.getConnection(URL, USER, PW);
            st = connection.createStatement();

            String sql;
            sql = "select * FROM BOARD;";

            ResultSet rs = st.executeQuery(sql);

            while (rs.next()) {
                String sqlRecipeProcess = rs.getString("BOARD_TITLE");
            }

            rs.close();
            st.close();
            connection.close();
        } catch (SQLException se1) {
            se1.printStackTrace();
        } catch (Exception ex) {
            ex.printStackTrace();
        } finally {
            try {
                if (st != null)
                    st.close();
            } catch (SQLException se2) {
            }
            try {
                if (connection != null)
                    connection.close();
            } catch (SQLException se) {
                se.printStackTrace();
            }
        }
    }
}
