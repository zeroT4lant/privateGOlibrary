package main

//ИНДЕКСЫ
//Индексы используются для ускорения процесса поиска данных.
//Индексы особенно полезны при выполнении операций сравнения "WHERE" "ORDER" И "JOIN"
//Но они занимают место в памяти и могут замедлить операции вставки/обновления/удаления/
//Одиночный индекс - create index 'название_индекса' on 'Таблица' ('нужное_поле1');
//Составной индекс - create index 'название_индекса' on 'Таблица' ('поле1','поле2','поле3');

//CREATE INDEX name ON table (column);
//CREATE INDEX name ON table USING index_type (column);

//Типы индексов
//B-Tree - ЭТО СБАЛАНСИРОВАНОЕ ДЕРЕВО, А НЕ БИНАРНОЕ
//B-Tree. Наиболее часто используемый тип индекса. Хорош для операций равенства и диапозонов.
//Hash. Эффективен для операций равенства и точечного поиска. Не поддерживает диапозонные запросы.
//GiST(Generalized Search Tree). Поддерживает балансировку дерева и может быть использован для различных типов данных.
//GIN(Generalized Inverted Index). Тип индекса, эффективный для значений, которые содержат несколько значений в одной колонке(массивы)

//Уровни изоляций Транзакций.
//1.Read Uncommitted - самый низкий уровень изоляций. Можно видеть грязные данные от других транзакций
//2.Read Committed - даёт доступ к данным, зафиксированным до начала транзакции.
//3.Repeatable Read - предотвращает неповторяемое чтение.
//4.Serializable - Предотвращает все виды конфликтов чтения и записи. Все транзакции выполняются в строгой последовательности.
//одна за другой.

//Денормализация - процесс уменьшения нормализации.Сокращаем количество таблиц/объединяем данные в одной таблице.
//Чтобы улучшить производительность.
//Причины использования денормализации.
//1.Улучшение производительности запросов.
//2.Упрощение запросов.
//3.Увеличение скорости чтения.

//Типы джоинов.
//Стандартный Inner JOIN - совпадения в двух таблицах.
//Left/Right JOIN - возвращает все строки из выбранной стороны таблицы и совпадающие строки из другой таблицы. При несовпадении NULL.
//Full JOIN - Возвращает строки, когда есть совпадение в одной из таблиц.Независимо от того есть ли значение в другой таблице.
//Если проще, то Inner+Left+Right JOIN.
//SELECT products.name, product_types.type_name
//FROM products JOIN product_types
//ON products.type_id = product_types.id

//Что можно сделать если запрос жирный и много JOIN`ов
//1.Индексация
//2.Уменьшение количество JOIN'ов.
//3.Использование подзапросов,вместо join'ов
//4.Оптимизируй запрос с помощью EXPLAIN
//5.Используй view, который хранит результат этого запроса.
//6.Вертикальное шардирование.

//ACID
//Atomicity(Атомарность)-Транзакция выполняется либо успешно, либо не выполняется совсем.
//Consistency(согласованность)-БД переходить из одного согласованного состояние в другое
//Isolation(Изолированность)-Использование уровней изоляции.
//Durability(Долговечность)-После завершения транзакции, данные сохраняются и остаются постоянными, усточивыми.Даже при сбое системы.

// Репликация - процесс копирования данных из одной БД(оригинал - мастер), в несколько копий других БД(реплики - слейв).
// В мастер БД записываем данные, из слейвов читаем. Это улучшает производительность.
// Если основная БД становится недоступной, то переключаемся на реплику без потери данных.
// Во время репликации обычно передаются следующие данные:
// 1.Транзакции
// 2.Состояние БД
// 3.Метаданные - информация о структуре БД и её объектах(таблицы,индексы,ограничения)

// TOAST - механизм, который позволяет хранить большие значения данных
// TOAST позволяет PostgreSQL эффективно работать с большими значениями данных, минимизируя использование дискового пространства и улучшая производительность.

// Для оптимизации запроса используй EXPAIN

// BIGINT VS UUID
// bigint bigint обычно используется, когда вам нужен простой, упорядоченный идентификатор. Он занимает меньше места (8 байтов) и обычно быстрее для операций сравнения и сортировки.
// uuid представляет собой универсальный уникальный идентификатор, который генерируется таким образом, что вероятность его дублирования крайне мала.
func main() {

}
