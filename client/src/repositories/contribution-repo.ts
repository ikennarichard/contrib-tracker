import { ApiClient } from '../api-client.js';
import { Contribution, CreateContribution } from '../types/contribution.js';

export class ContributionRepository {
  constructor(private readonly apiClient: ApiClient) {}

  async add(data: CreateContribution): Promise<Contribution> {
    return this.apiClient.addContribution(data);
  }

  async list(repository?: string): Promise<Contribution[]> {
    return this.apiClient.listContributions(repository);
  }
}