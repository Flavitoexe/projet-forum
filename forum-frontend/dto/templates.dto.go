package dto

// DTO pour index.html
type TempIndexData struct {
	ThreadsData    TempThreadsData
	CategoriesData TempCategoriesData
	UserData       TempUserData
}

// DTO pour thread.html
type TempThreadDetailsData struct {
	ThreadData     TempThreadData
	MessagesData   TempMessagesData
	CategoriesData TempCategoriesData
	UserData       TempUserData
}

// DTO pour category.html
type TempCategoryDetailsData struct {
	ThreadsData    TempThreadsData
	CategoryData   TempCategoryData
	CategoriesData TempCategoriesData
	UserData       TempUserData
}

type TempCategoriesDetailsData struct {
	CategoriesData TempCategoriesData
	UserData       TempUserData
}

type TempErrorDetailsData struct {
	ErrorData DtoError
	UserData  TempUserData
}
