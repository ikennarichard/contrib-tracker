#!/usr/bin/env tsx
import { ApiClient } from './src/api-client.js';
import { ContributionRepository } from './src/repositories/contribution-repo.js';
import { ContributionService } from './src/services/contribution-service.js';

const apiClient = new ApiClient();
const repository = new ContributionRepository(apiClient);
const service = new ContributionService(repository);

async function main() {
  const [command, ...args] = process.argv.slice(2);

  try {
    if (command === 'add') {
      const title = getFlag(args, ['--title', '-t']);
      const repo = getFlag(args, ['--repo', '-r', '--repository']);
      const url = getFlag(args, ['--url', '-u']);
      const date = getFlag(args, ['--date', '-d']);

      if (!title || !repo) {
        console.error('Usage: tsx contrib.ts add --title "Title" --repo "owner/repo" [--url "https://..."] [--date "YYYY-MM-DD"]');
        process.exit(1);
      }

      const result = await service.logContribution(title, repo, url, date);
      console.log('Contribution added successfully:');
      console.dir(result, { depth: null });

    } else if (command === 'list') {
      const repoFilter = getFlag(args, ['--repo', '-r', '--repository']);
      const contributions = await service.getContributions(repoFilter);

      console.log(`📋 Contributions${repoFilter ? ` for ${repoFilter}` : ''}:`);
      if (contributions.length === 0) {
        console.log('No contributions found.');
      } else {
        contributions.forEach(c => {
          console.log(`• [${c.date}] ${c.title} (${c.repository})${c.url ? ` → ${c.url}` : ''}`);
        });
      }
    } else {
      console.log(`
Usage:
  tsx contrib.ts add --title "..." --repo "..." [--url "..."] [--date "YYYY-MM-DD"]
  tsx contrib.ts list [--repo "..."]
      `);
    }
  } catch (error: any) {
    console.error('Error:', error.message);
    process.exit(1);
  }
}

function getFlag(args: string[], keys: string[]): string | undefined {
  for (let i = 0; i < args.length; i++) {
    if (keys.includes(args[i])) {
      return args[i + 1];
    }
  }
  return undefined;
}

main();