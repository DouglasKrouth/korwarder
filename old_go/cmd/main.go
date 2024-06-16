// User entry point for `korwarder`. Responsible for starting cli and all related funcs.
package main

import (
// "log"
// "os/exec"
//    "fmt"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
)

// func pf(port int) {
//     p := fmt.Sprintf("27017:%d", port)
// 	cmd := exec.Command("kubectl", "port-forward", "mongo-7d96cb4cf-7t7q4", p)
// 	err := cmd.Start()
// 	if err != nil {
//         fmt.Println("PIGGGGGGU")
// 		log.Fatal(err)
// 	}
// 	log.Printf("Waiting for command to finish...")
// 	err = cmd.Wait()
// 	log.Printf("Command finished with error: %v", err)
//
// }

func main() {
    PortForward("kubectl do something")
    zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
    log.Print("hello world")


    // pf(27017)
    // go pf(27087)

}
