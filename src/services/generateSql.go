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
        sql := fmt.Sprintf("INSERT INTO hard_skills.tbl_test (id_test, name, duration, locked, state, tags) VALUES ('%s', '%s', %d, %t, '%s', ARRAY%s::text[]);\n",
            test.ID,
            strings.ReplaceAll(test.Name, "'", "''"),
            test.Duration,
            test.Locked,
            test.State,
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
                candidate.ID,
                candidate.Email,
                candidate.Full_name,
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
                candidate.Test_id,
                candidate.Email,
                candidate.ID,
                candidate.Score,
                candidate.Attempt_starttime,
                candidate.Attempt_endtime,
                candidate.Invited_on,
                candidate.Pdf_url,
                scoreTagsSplitJSON,
                questionsJSON,
                candidate.Feedback,
                candidate.Percentage_score,
            )
            candidateTestFile.WriteString(candidateTestSQL)

            var tags []string
            for tag := range candidate.Score_tags_split {
                if !contains(tags, tag) {
                    tags = append(tags, tag)

                    tagSQL := fmt.Sprintf("INSERT INTO hard_skills.tbl_tag (tag_name) VALUES ('%s');\n", tag)
                    tagFile.WriteString(tagSQL)

                    crossTagSQL := fmt.Sprintf("INSERT INTO hard_skills.tbl_cross_tag (attempt_id, tag_id, score) VALUES ('%s', (SELECT id_tag FROM hard_skills.tbl_tag WHERE tag_name = '%s' LIMIT 1), '%s');\n", candidate.ID, tag, candidate.Score_tags_split[tag])
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