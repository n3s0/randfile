package main

import (
    "fmt"
    "os"
    "errors"
    "path/filepath"
    "math/rand/v2"
    
    "github.com/spf13/cobra"
)

var path string

var rootCmd = &cobra.Command{
    Use:   "randfile",
    Short: "randfile picks a random file from a directory path",
    Long: `randfile picks a random file from a directory path.`,
    Run: func(cmd *cobra.Command, args []string) {
        if path != "" {
            filePaths, err := directoryPathContents(path)
            if err != nil {
                fmt.Println(err)
                os.Exit(1)
            }
            
            chosenPath, err := randFile(filePaths)
            if err != nil {
                fmt.Println(err)
                os.Exit(1)
            }
            
            fmt.Println(chosenPath)
        }
    },
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version number of randfile",
  Long:  `Print the version number of randfile`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("randfile v0.1")
  },
}

func main() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func init() {
    rootCmd.AddCommand(versionCmd)
    rootCmd.Flags().StringVarP(&path, "path", "p", "", "directory path to read from")

}

func directoryPathContents(dirPath string) (filePaths []string, err error) {
    absPath, err := filepath.Abs(dirPath)
    if os.IsNotExist(err) {
        return nil, errors.New("Path does not exist!")
    } else if err != nil {
        fmt.Println(err)
    }

    pathContents, err := os.ReadDir(absPath)
    if err != nil {
        fmt.Println(err)
    }

    for _, pathObject := range pathContents {
        if !pathObject.IsDir() {
            fullPath := filepath.Join(absPath, pathObject.Name())
            filePaths = append(filePaths, fullPath)
        }
    }

    return filePaths, nil
}

func randFile(filePaths []string) (filePath string, err error) {
    count := len(filePaths)

    if count <= 0 {
        return "", errors.New("Number of files and directories is 0.")
    }
    
    filePathIdx := rand.IntN(count)
    
    filePath = filePaths[filePathIdx]
    
    return filePath, nil
}

