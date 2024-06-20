import type {
  ViewMatchDetailsV2,
  ViewMatchMMRCalculationDetails,
  ViewMatchTeamV2,
} from '$api';

export const movePlayerToMember1 = (
  match: ViewMatchDetailsV2,
  playerId: number
): ViewMatchDetailsV2 => {
  // Player is already member 1
  if (match.team1.member1 === playerId) {
    return match;
  }

  if (match.team1.member2 === playerId) {
    return {
      ...match,
      team1: movePlayerInTeam(match.team1, { from: 'member2', to: 'member1' }),
      mmrCalculations: flipMembersOfTeamInMMRCalculation(
        match.mmrCalculations,
        'team1'
      ),
    };
  }

  if (match.team2.member1 === playerId) {
    return {
      ...match,
      team1: match.team2,
      team2: match.team1,
      mmrCalculations: flipTeamsInMMRCalculation(match.mmrCalculations),
    };
  }

  return {
    ...match,
    team1: movePlayerInTeam(match.team2, { from: 'member2', to: 'member1' }),
    team2: match.team1,
    mmrCalculations: flipMembersOfTeamInMMRCalculation(
      flipTeamsInMMRCalculation(match.mmrCalculations),
      'team1'
    ),
  };
};

const movePlayerInTeam = (
  team: ViewMatchTeamV2,
  { from, to }: { from: 'member1' | 'member2'; to: 'member1' | 'member2' }
) => {
  return {
    ...team,
    [to]: team[from],
    [from]: team[to],
  };
};

const flipMembersOfTeamInMMRCalculation = (
  mmrCalculations: ViewMatchMMRCalculationDetails | undefined,
  team: 'team1' | 'team2'
): ViewMatchMMRCalculationDetails | undefined => {
  if (mmrCalculations == null) {
    return mmrCalculations;
  }
  return {
    ...mmrCalculations,
    [team]: {
      player1MMRDelta: mmrCalculations[team].player2MMRDelta,
      player2MMRDelta: mmrCalculations[team].player1MMRDelta,
    },
  };
};

const flipTeamsInMMRCalculation = (
  mmrCalculations: ViewMatchMMRCalculationDetails | undefined
): ViewMatchMMRCalculationDetails | undefined => {
  if (mmrCalculations == null) {
    return mmrCalculations;
  }
  return {
    team1: mmrCalculations.team2,
    team2: mmrCalculations.team1,
  };
};
