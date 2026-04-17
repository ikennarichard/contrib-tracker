import { execSync } from "node:child_process";

function run() {
  const title = "ts made this contribution";
  const repo = "demo/repo";
  const url = "https://github.com/demo/repo";

  try {
    execSync(
      `go run main.go add --title "${title}" --repo "${repo}" --url "${url}"`,
      { stdio: "inherit" },
    );
  } catch (err) {
    console.error("Failed to log contribution");
    process.exit(1);
  }
}

run();
