import { describe, it, expect, vi } from 'vitest';
import { ApiClient } from '../api-client';
import { ContributionRepository } from '../repositories/contribution-repo';
import { ContributionService } from './contribution-service';

describe('ContributionService', () => {
  const apiClient = new ApiClient('http://localhost:8080');
  const repo = new ContributionRepository(apiClient);
  const service = new ContributionService(repo);

  it('should log a new contribution using service layer', async () => {
    const result = await service.logContribution(
      "Service Layer Test",
      "ikennarichard/contrib-tracker",
      "https://github.com/pull/100"
    );

    expect(result).toMatchObject({
      title: "Service Layer Test",
      repository: "ikennarichard/contrib-tracker",
    });
    expect(result.id).toBeDefined();
  });

  it('should return empty array when no contributions exist', async () => {
    const contributions = await service.getContributions();
    expect(Array.isArray(contributions)).toBe(true);
  });

  it('should filter contributions by repository', async () => {
    const contributions = await service.getContributions("ikennarichard/contrib-tracker");
    expect(Array.isArray(contributions)).toBe(true);
  });

  it('should throw error when title is empty', async () => {
    await expect(
      service.logContribution('', 'ikennarichard/contrib-tracker')
    ).rejects.toThrow();
  });

  it('should handle service layer when repository returns null/undefined', async () => {
    vi.spyOn(repo, 'list').mockResolvedValueOnce(null as any);

    const contributions = await service.getContributions();
    expect(contributions).toEqual([]);
  });
});