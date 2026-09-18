# `tv-episode-groups`

TV episode group details.

| action    | example                                                              | output file                                                  |
|-----------|-----------------------------------------------------------------------|-----------------------------------------------------------------|
| `details` | `tmdb-sync tv-episode-groups -a details -i 5acf98d40e0a26346a0389c3` | `tv-episode-groups_details_id-5acf98d40e0a26346a0389c3.json` |

Notes:
- `-i <episode_group_id>` is required.
- This is distinct from `tv -a episode-groups -i <series_id>`, which lists a series' episode
  groups; this module fetches the full details (including grouped episodes) for a single episode
  group by its own id.
