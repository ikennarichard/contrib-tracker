import { Contribution, CreateContribution } from './types/contribution.js';

export class ApiClient {
  private readonly baseUrl: string;

  constructor(baseUrl: string = process.env.API_BASE_URL ?? 'http://localhost:8080') {
    this.baseUrl = baseUrl.replace(/\/$/, '');
  }

  async addContribution(data: CreateContribution): Promise<Contribution> {
    const response = await fetch(`${this.baseUrl}/api/contributions`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data),
    });

    if (!response.ok) {
      const errorText = await response.text();
      throw new Error(`Failed to add contribution: ${response.status} ${errorText}`);
    }

    return response.json();
  }

  async listContributions(repository?: string): Promise<Contribution[]> {
    const url = new URL(`${this.baseUrl}/api/contributions`);
    if (repository) {
      url.searchParams.set('repository', repository);
    }

    const response = await fetch(url.toString());
    if (!response.ok) {
      throw new Error(`Failed to fetch contributions: ${response.status}`);
    }

    return response.json();
  }
}