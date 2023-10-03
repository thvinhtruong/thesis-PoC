package mock

func MockModuleData(course string) []string {
	csTopics := []string{"Linear Algebra", "Calculus", "Discrete Math",
		"Introduction to Computer Science", "C/C++", "Python",
		"Web Programing", "Data Science", "Projects",
		"Current Topics In Computer Science", "Computer Network",
		"Software Design", "Law and Data Protection", "Artificial Intelligence",
		"Software Development", "Java", "Operating System", "Distributed System",
		"Internship", "Thesis"}

	faTopics := []string{"OMath", "OWin", "OStats",
		"Data Analytics", "OLaw", "Basic Accounting",
		"Finance", "Business Administration", "Marco Economics",
		"Micro Economics", "Advanced Accounting", "Group Accounting",
		"Marketing", "Research Methodology", "Python",
		"Global Finance", "IT in Business", "Communication Skills",
		"Internship", "Thesis"}

	switch course {
	case "CS":
		return csTopics
	case "FA":
		return faTopics
	case "BA":
		return faTopics
	default:
		return nil
	}
}
