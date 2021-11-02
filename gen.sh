ROOT_PACKAGE="/Users/zouchunxu/web/docker/wwwroot/go/my-operator"

CUSTOM_RESOURCE_NAME="samplecrd"


CUSTOM_RESOURCE_VERSION="v1"


./generate-groups.sh all "$ROOT_PACKAGE/pkg/client" "$ROOT_PACKAGE/pkg/apis" "$CUSTOM_RESOURCE_NAME:$CUSTOM_RESOURCE_VERSION"