import org.junit.Test;

import java.sql.Connection;
import java.sql.DriverManager;

public class MySQLConnectionTest {
    
    private static final String DRIVER = "com.mysql.jdbc.Driver";
    private static final String URL = "jdbc:mysql://webdbinstance.cc1j8x3t3gto.ap-northeast-2.rds.amazonaws.com/webdb";
    private static final String USER = "dev";
    private static final String PW = "qrwgDGCGnMPGHtKqQhHnZYmhXfFxHyaJsQuAxnGEv";

    @Test
    public void testConnection() throws Exception {
        Class.forName(DRIVER);
        try(Connection conn = DriverManager.getConnection(URL, USER, PW)) {
            System.out.println(conn);
        } catch(Exception e) {
            e.printStackTrace();
        }
    }
}
