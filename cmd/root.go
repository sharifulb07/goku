package cmd

import (
	"fmt"
	"os"

	"github.com/sharifulb07/goku/internal/db"
	"github.com/sharifulb07/goku/internal/parser"
	"github.com/sharifulb07/goku/internal/repository"
	"github.com/spf13/cobra"
)

var (
	inputFile  string
	dbConn     string
	resourceID int
)

// ====================== Persistent PreRun ======================
func initDB(cmd *cobra.Command, args []string) error {
	if cmd.Use == "save" || cmd.Use == "list" || cmd.Use == "update" ||
		(cmd.Parent() != nil && cmd.Parent().Use == "delete") {
		return db.Init(dbConn)
	}
	return nil
}

// ====================== Run Functions ======================
func saveRun(cmd *cobra.Command, args []string) error {
	repo := repository.NewResourceRepo()
	resource, err := parser.ReadResourceJSON(inputFile)
	if err != nil {
		return err
	}

	if err := repo.Create(resource); err != nil {
		return err
	}

	fmt.Printf("✅ Resource '%s' saved successfully! (ID: %d)\n", resource.Name, resource.ID)
	return nil
}

func listRun(cmd *cobra.Command, args []string) error {
	repo := repository.NewResourceRepo()
	resources, err := repo.List()
	if err != nil {
		return err
	}

	if len(resources) == 0 {
		fmt.Println("No resources found.")
		return nil
	}

	for _, r := range resources {
		fmt.Printf("%d | %s | %s | %s | %s\n", r.ID, r.Name, r.Role, r.Stack, r.Status)
	}
	return nil
}

func updateRun(cmd *cobra.Command, args []string) error {
	if resourceID == 0 {
		return fmt.Errorf("resource ID is required (--id)")
	}
	repo := repository.NewResourceRepo()
	resource, err := parser.ReadResourceJSON(inputFile)
	if err != nil {
		return err
	}

	if err := repo.Update(resourceID, resource); err != nil {
		return err
	}

	fmt.Printf("✅ Resource ID %d updated successfully!\n", resourceID)
	return nil
}

func softDeleteRun(cmd *cobra.Command, args []string) error {
	if resourceID == 0 {
		return fmt.Errorf("resource ID is required (--id)")
	}
	repo := repository.NewResourceRepo()
	return repo.Delete(resourceID)
}

func hardDeleteRun(cmd *cobra.Command, args []string) error {
	if resourceID == 0 {
		return fmt.Errorf("resource ID is required (--id)")
	}
	repo := repository.NewResourceRepo()
	return repo.HardDelete(resourceID)
}

// ====================== Commands ======================
var rootCmd = &cobra.Command{
	Use:   "goku",
	Short: "Goku CLI - JSON/YAML Converter & CRUD Tool",
}

var saveCmd = &cobra.Command{Use: "save", Short: "Save resource", RunE: saveRun}
var listCmd = &cobra.Command{Use: "list", Short: "List resources", RunE: listRun}
var updateCmd = &cobra.Command{Use: "update", Short: "Update resource", RunE: updateRun}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a resource",
}

var hardDeleteCmd = &cobra.Command{
	Use:   "hard",
	Short: "Permanently (hard) delete a resource",
	RunE:  hardDeleteRun,
}

// ====================== Execute ======================
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&dbConn, "db", "postgres://postgres:5511@localhost:5432/goku?sslmode=disable", "PostgreSQL connection string")
	rootCmd.PersistentPreRunE = initDB

	// Add main commands
	rootCmd.AddCommand(saveCmd, listCmd, updateCmd, deleteCmd)

	// Add hard delete as sub-command under delete
	deleteCmd.AddCommand(hardDeleteCmd)

	// Flags
	saveCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Input JSON file")
	saveCmd.MarkFlagRequired("input")

	updateCmd.Flags().StringVarP(&inputFile, "input", "i", "", "JSON file with update data")
	updateCmd.Flags().IntVar(&resourceID, "id", 0, "Resource ID")
	updateCmd.MarkFlagRequired("input")
	updateCmd.MarkFlagRequired("id")

	deleteCmd.Flags().IntVar(&resourceID, "id", 0, "Resource ID")
	deleteCmd.MarkFlagRequired("id")

	hardDeleteCmd.Flags().IntVar(&resourceID, "id", 0, "Resource ID")
	hardDeleteCmd.MarkFlagRequired("id")
}