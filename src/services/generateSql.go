package services

import (
	"encoding/json"
	"fmt"
	"hacker-rank-lambda/src/structures"
	"hacker-rank-lambda/src/utils"
	"os"
	"strconv"

	"strings"
)

func GenerateTestSQL(tests []structures.Tests) error {
    file, err := os.Create("Tests.sh")
    if err != nil {
        return err
    }
    defer file.Close()
    

    for _, test := range tests {
        name := strings.ReplaceAll(test.Name, "'", "")
        sql := fmt.Sprintf("INSERT INTO hard_skills.tbl_test (id_test, name, duration, locked, state, tags) VALUES ('%s', '%s', %d, %t, '%s', ARRAY%s::text[]);\n",
            test.ID,
            name,
            test.Duration,
            test.Locked,
            strings.ReplaceAll(test.State, "'", ""),
            utils.FormatArray(test.Tags),
        )
        file.WriteString(sql)
    }

    fmt.Println("Shell script generated successfully!")
    return nil
}

func GenerateCandidateSQL(candidates []structures.CandidateResponse) error {
    file, err := os.Create("Candidate.sh")
    if err != nil {
        return err
    }
    defer file.Close()

    for _, candidateResponse := range candidates {
        for _, candidate := range candidateResponse.Data {
            sql := fmt.Sprintf("INSERT INTO hard_skills.tbl_candidate (id_candidate, email, full_name) VALUES ('%s', '%s', '%s');\n",
                strings.ReplaceAll(candidate.ID, "'", ""),
                strings.ReplaceAll(candidate.Email, "'", ""),
                strings.ReplaceAll(candidate.Full_name, "'", ""),
            )
            file.WriteString(sql)
        }
    }

    fmt.Println("Shell script generated successfully!")
    return nil
}

func GenerateCandidateTestSQL(candidates []structures.CandidateResponse) error {
    candidateTestFile, err := os.Create("Candidate_test.sh")
    if err != nil {
        return err
    }
    defer candidateTestFile.Close()
    tagMap := make(map[string]bool)
    tagFile, err := os.Create("Tag.sh")
    if err != nil {
        return err
    }
    defer tagFile.Close()

    crossTagFile, err := os.Create("Cross_tag.sh")
    if err != nil {
        return err
    }
    defer crossTagFile.Close()

    for _, candidateResponse := range candidates {
        for _, candidate := range candidateResponse.Data {
            scoreTagsSplitJSON, _ := json.Marshal(candidate.Score_tags_split)
            questionsFloat := make(map[string]float64)
            for k, v := range candidate.Questions {
                f, _ := strconv.ParseFloat(fmt.Sprintf("%v", v), 64)
                questionsFloat[k] = f
            }

            questionsJSON, _ := json.Marshal(questionsFloat)

            candidateTestSQL := fmt.Sprintf("INSERT INTO hard_skills.tbl_candidate_test (id_test, email, attempt_id, score, attempt_starttime, attempt_endtime, invited_on, pdf_url, score_tags_split, questions, feedback, percentage_score) VALUES ('%s', '%s', '%s', %f, '%s', '%s', '%s', '%s', '%s', '%s', '%s', %f);\n",
                strings.ReplaceAll(candidate.Test_id, "'", ""),
                strings.ReplaceAll(candidate.Email, "'", ""),
                strings.ReplaceAll(candidate.ID, "'", ""),
                candidate.Score,
                candidate.Attempt_starttime,
                candidate.Attempt_endtime,
                candidate.Invited_on,
                candidate.Pdf_url,
                scoreTagsSplitJSON,
                questionsJSON,
                strings.ReplaceAll(candidate.Feedback, "\n", " "),
                candidate.Percentage_score,
            )
            candidateTestFile.WriteString(candidateTestSQL)

            var tags []string
            for tag := range candidate.Score_tags_split {
                if _, ok := tagMap[tag]; !ok {
                    tagMap[tag] = true

                    tagSQL := fmt.Sprintf("INSERT INTO hard_skills.tbl_tag (id, tag_name) SELECT COALESCE(MAX(CAST(id AS INTEGER)), 0) + 1, '%s' FROM hard_skills.tbl_tag WHERE NOT EXISTS (SELECT 1 FROM hard_skills.tbl_tag WHERE tag_name = '%s');\n", tag, tag)
                    tagFile.WriteString(tagSQL)
                }
                if !contains(tags, tag) {
                    tags = append(tags, tag)

                    crossTagSQL := fmt.Sprintf("INSERT INTO hard_skills.tbl_cross_tag (attempt_id, tag_id, score) VALUES ('%s', (SELECT id FROM hard_skills.tbl_tag WHERE tag_name = '%s' LIMIT 1), '%s');\n", candidate.ID, tag, candidate.Score_tags_split[tag])
                    crossTagFile.WriteString(crossTagSQL)
                }
            }
        }
    }

    fmt.Println("Shell scripts generated successfully!")
    return nil
}

func contains(slice []string, str string) bool {
    for _, s := range slice {
        if s == str {
            return true
        }
    }
    return false
}