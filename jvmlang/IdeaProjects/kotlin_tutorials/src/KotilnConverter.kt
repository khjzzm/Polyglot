import org.junit.Test

import java.sql.DriverManager

class MySQLConnectionTest {

    @Test
    @Throws(Exception::class)
    fun testConnection() {
        Class.forName(DRIVER)
        try {
            DriverManager.getConnection(URL, USER, PW).use { conn -> println(conn) }
        } catch (e: Exception) {
            e.printStackTrace()
        }

    }

    companion object {
        private val DRIVER = "com.mysql.jdbc.Driver"
        private val URL = "jdbc:mysql://webdbinstance.cc1j8x3t3gto.ap-northeast-2.rds.amazonaws.com/webdb"
        private val USER = "dev"
        private val PW = "qrwgDGCGnMPGHtKqQhHnZYmhXfFxHyaJsQuAxnGEv"
    }
}
