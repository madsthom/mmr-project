import type { ReposLeaderboardEntry } from '../../../api';

export interface LeaderboardEntry extends ReposLeaderboardEntry {
  rank: number;
}
