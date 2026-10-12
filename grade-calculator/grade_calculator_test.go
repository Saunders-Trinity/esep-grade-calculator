package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 95, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 91, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetFinalGradeThresholds(t *testing.T) {
	tests := []struct {
		grade    int
		expected string
	}{
		{90, "A"},
		{80, "B"},
		{70, "C"},
		{60, "D"},
		{59, "F"},
	}

	for _, test := range tests {
		calculator := NewGradeCalculator()
		calculator.AddGrade("assignment", test.grade, Assignment)
		calculator.AddGrade("exam", test.grade, Exam)
		calculator.AddGrade("essay", test.grade, Essay)

		if actual := calculator.GetFinalGrade(); actual != test.expected {
			t.Errorf("grade %d: expected %q; got %q", test.grade, test.expected, actual)
		}
	}
}

func TestGetFinalGradeWithoutGrades(t *testing.T) {
	if actual := NewGradeCalculator().GetFinalGrade(); actual != "F" {
		t.Errorf("expected empty calculator to return F; got %q", actual)
	}
}

func TestGradeTypeString(t *testing.T) {
	if actual := Exam.String(); actual != "exam" {
		t.Errorf("expected Exam.String() to return %q; got %q", "exam", actual)
	}
}
