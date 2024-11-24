package database

import (
	"fmt"
	"typehero_server/models"
)



func (db *Database) GetLeaderboard(req models.LeaderboardRequest) ([]models.Result, error) {
    var results []models.Result
    err :=db.Raw(fmt.Sprintf(`
WITH ranked_results AS (
    SELECT
        *,
        MAX(wpm) AS max_wpm
    FROM results
    WHERE mode = '%s'
      AND word_amount = %d
      AND language = '%s'
    GROUP BY user_id
)
SELECT
        *,
    RANK() OVER (ORDER BY max_wpm DESC) AS rank,
    max_wpm
FROM ranked_results
ORDER BY rank
LIMIT 25 OFFSET 25 * (%d - 1);
        `, req.Mode, req.WordAmount, req.Language,  req.Page)).Find(&results).Error
    if err != nil {
        return nil, err
    }
    return results, err
}
