package sems

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"regexp"
	"strconv"
	"time"
)

const semsAddr = "127.0.0.1"
const semsPort = "5040"

func openSemsStatsConn() net.Conn {
	conn, err := net.DialTimeout("udp", fmt.Sprintf("%s:%s", semsAddr, semsPort), 500*time.Millisecond)
	if err != nil {
		fmt.Printf("err opening sems connection: %+v", err)
		return nil
	}
	return conn
}

func sendSemsStats(cmd string, conn net.Conn) (string, error) {
	if conn == nil {
		return "", errors.New("Conn is nil")
	}
	defer conn.Close()

	fmt.Printf("Writing to connection: %s\n", cmd)
	fmt.Fprintf(conn, cmd+"\n")
	status, err := bufio.NewReader(conn).ReadString('\n')

	return status, err
}

// GetActiveCallCount returns the number of ongoing calls in SEMS
func GetActiveCallCount() (int, error) {
	resp, err := sendSemsStats("calls", openSemsStatsConn())
	if err != nil {
		fmt.Printf("SEMS - error: %v\n", err)
		return 0, err
	}
	fmt.Printf("SEMS - cmdActiveCallCount response: %s\n", resp)

	// parse response to int
	re := regexp.MustCompile(`Active calls: (\d*)`)
	matches := re.FindStringSubmatch(resp)
	if len(matches) > 1 {
		calls, err := strconv.Atoi(matches[1])
		if err != nil {
			fmt.Printf("SEMS - GetActiveCallCount: could not convert response to int: %v\n", err)
			return 0, err
		}
		return calls, nil
	}
	return 0, errors.New("SEMS - Unexpected response to command")
}
