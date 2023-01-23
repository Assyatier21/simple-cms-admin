echo "==generating mockfile for repository=="
mockgen -source=./internal/repository/postgres/init.go -destination=./mock/repository/postgres/init.go
echo "==mockfile for repository generated=="

echo "==generating mockfile for usecase=="
mockgen -source=./internal/usecase/init.go -destination=./mock/usecase/init.go
echo "==mockfile for usecase generated=="

echo "==generating mockfile for api handler=="
mockgen -source=./internal/delivery/api/init.go -destination=./mock/delivery/api/init.go
echo "==mockfile for api handler generated==" 
