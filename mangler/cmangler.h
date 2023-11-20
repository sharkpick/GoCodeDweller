#ifndef __C_MANGLER_H__
#define __C_MANGLER_H__

#ifdef __cplusplus
extern "C"
{
#endif

    typedef void *Mangler;

    Mangler NewMangler();
    void DestroyMangler(Mangler mangler);
    unsigned char Encrypt(Mangler mangler, unsigned char c);
    unsigned char Decrypt(Mangler mangler, unsigned char c);

#ifdef __cplusplus
}
#endif

#endif // __C_MANGLER_H__
