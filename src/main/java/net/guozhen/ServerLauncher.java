package net.guozhen;

import net.guozhen.server.NettyServer;

/**
 * java -classpath target/netty-1.0.0-SNAPSHOT.jar com.guozhen.java.netty.ServerLauncher
 *
 * 访问 http://127.0.0.1:10089/
 */
public class ServerLauncher {
    public static void main(String[] args) {
        NettyServer.start(10089);
    }
}
