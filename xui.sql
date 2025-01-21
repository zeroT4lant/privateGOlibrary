SELECT DISTINCT 
user.id,
purchase.sku
from user
join purchase on user.id = purchase.user_id
left join ban_list on user.id = ban_list.user_id
where purchase.date < ban_list.date_from or ban_list.date_from is null
order by user.id,sku