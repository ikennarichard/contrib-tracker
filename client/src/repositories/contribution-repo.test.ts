import { describe, it, expect } from 'vitest';
import { ApiClient } from '../api-client';
import { ContributionRepository } from './contribution-repo';

describe('ContributionRepository', () => {
  const apiClient = new ApiClient('http://localhost:8080');
  const repository = new ContributionRepository(apiClient);

  it('should add contribution through repository', async () => {
    const result = await repository.add({
      title: "Repository Pattern Test",
      repository: "ikennarichard/contrib-tracker",
      date: "2026-04-18"
    });

    expect(result.id).toBeDefined();
    expect(result.title).toBe("Repository Pattern Test");
  });

  it('should list contributions through repository', async () => {
    const contributions = await repository.list();

    expect(Array.isArray(contributions)).toBe(true);
  });
});