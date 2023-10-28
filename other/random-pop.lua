 -- 定义一个函数，用于从一个列表中随机选择 n 个元素，并返回一个包含这些元素的表
local function random_select(list, n)
    -- 创建一个空表，用于存储结果
    local result = {}
    -- 如果列表为空或者 n 小于等于 0，直接返回空表
    if #list == 0 or n <= 0 then
        return result
    end
    -- 如果 n 大于等于列表的长度，直接返回整个列表
    if n >= #list then
        return list
    end
    -- 创建一个表，用于存储已经选择过的索引，避免重复选择
    local selected = {}
    -- 循环 n 次，每次随机选择一个元素
    for i = 1, n do
        -- 生成一个随机数，范围是 1 到列表的长度
        local index = math.random(1, #list)
        -- 检查这个索引是否已经被选择过，如果是，就重新生成一个随机数，直到找到一个没有被选择过的索引
        while selected[index] do
            index = math.random(1, #list)
        end
        -- 把这个索引对应的元素加入到结果表中
        table.insert(result, list[index])
        -- 把这个索引标记为已经被选择过
        selected[index] = true
    end
    -- 返回结果表
    return result
end

-- 获取传入的 key 和 count 参数，key 是要操作的列表的键名，count 是要弹出的元素个数
local key = KEYS[1]
local count = tonumber(ARGV[1])

-- 获取列表中所有的元素，存储在一个表中
local list = redis.call('lrange', key, 0, -1)

-- 调用上面定义的函数，从列表中随机选择 count 个元素，并存储在另一个表中
local selected = random_select(list, count)

-- 遍历被选择的元素表，从原来的列表中删除这些元素（注意：这里假设列表中没有重复的元素）
for i = 1, #selected do
    redis.call('lrem', key, 0, selected[i])
end

-- 返回被选择的元素表作为结果（注意：这里返回的是一个 Lua 表，不是 Redis 列表）
return selected