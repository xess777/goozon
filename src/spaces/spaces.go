package spaces

// RemoveExtraSpaces удаляет лишние пробелы в строке, представленной срезом rune.
// Возвращает указатель на модифицированный срез без лишних пробелов.
// Реализована с помощью Inplace алгоритма со сложность O(N) по времени и O(1) по памяти.
func RemoveExtraSpaces(sPtr *[]rune) *[]rune {
	s := *sPtr
	length := len(s)
	// i указывает на позицию заполнения итоговой строки.
	// j указывает на позицию символа в исходной строке.
	i, j := 0, 0
	// Флаг, который показывает, был ли предыдущий символ пробелом.
	isPreviousSpace := false

	// Увеличиваем счетчик j для пробелов вначале строки.
	for j < length && s[j] == ' ' {
		j++
	}

	// Основной цикл алгоритма, копируем символ из j в i, если это не пробел.
	// Игнорируем пробел, если предыдущий символ был пробелом.
	for ; j < length; j++ {
		if s[j] != ' ' {
			s[i] = s[j]
			isPreviousSpace = false
			i++
		} else if !isPreviousSpace {
			s[i] = s[j]
			isPreviousSpace = true
			i++
		}
	}

	// Стыдно за такой корявый код в конце, думаю можно было и лучше оформить.
	// Здесь разруливаем граничные кейсы, когда последний символ был пробелом.
	if i <= 1 || s[i-1] != ' ' {
		r := s[:i]
		return &r
	} else {
		r := s[:i-1]
		return &r
	}
}
